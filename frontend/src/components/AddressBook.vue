<template>
  <v-container class="fill-height outlined-row">
    <!--
    <v-responsive class="align-centerfill-height mx-auto" max-width="900">
      -->
    <v-row>
      <v-col cols="4">
        <v-list v-model:selected="selectedPerson">
          <v-list-subheader>Persons</v-list-subheader>
          <v-list-item
            v-for="people in book.people"
            :key="people.id"
            :title="people.name"
            :value="people"
          ></v-list-item>
        </v-list>
      </v-col>
      <v-divider class="ms-3" inset vertical></v-divider>
      <v-col cols="7">
        <v-row>
          <v-col cols="12">
            <v-text-field label="Name" v-model="personName"></v-text-field>
          </v-col>
          <v-col cols="12">
            <v-text-field label="Email" v-model="personEmail"></v-text-field>
          </v-col>
          <v-col cols="5">
            <v-list
              lines="one"
              style="height: 200px"
              v-model:selected="phoneInfo"
            >
              <v-list-subheader>Phone numbers</v-list-subheader>
              <!-- personPhones -->
              <v-list-item
                v-for="n in personPhones"
                :key="n.number"
                :title="n.number"
                :value="n"
              ></v-list-item>
            </v-list>
          </v-col>
          <v-col cols="7">
            <v-row>
              <v-text-field label="phone" v-model="phoneNumber"></v-text-field>
            </v-row>
            <v-row>
              <v-col cols="6">
                <v-radio-group v-model="phoneType">
                  <v-radio
                    label="移动电话"
                    :value="pb.PhoneType.MOBILE"
                  ></v-radio>
                  <v-radio
                    label="办公电话"
                    :value="pb.PhoneType.WORK"
                  ></v-radio>
                  <v-radio
                    label="家庭电话"
                    :value="pb.PhoneType.HOME"
                  ></v-radio>
                  <v-radio
                    label="其它"
                    :value="pb.PhoneType.UNSPECIFIED"
                  ></v-radio>
                </v-radio-group>
              </v-col>
              <v-col cols="6" align-self="end">
                <v-row>
                  <v-col cols="12">
                    <v-btn @click="savePhone">保存</v-btn>
                  </v-col>
                  <v-col cols="12">
                    <v-btn @click="delPhone">删除</v-btn>
                  </v-col>
                </v-row>
              </v-col>
            </v-row>
          </v-col>
        </v-row>
        <v-row>
          <v-divider class="ms-3" inset></v-divider>
        </v-row>
        <v-row>
          <v-col cols="2">
            <v-btn @click="savePerson">保存</v-btn>
          </v-col>
          <v-col cols="2">
            <v-btn @click="delPerson">删除</v-btn>
          </v-col>
        </v-row>
      </v-col>
    </v-row>
    <!--
    </v-responsive>
    -->
  </v-container>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";

import * as nats_core from "@nats-io/nats-core";
import {
  create,
  fromBinary,
  isMessage,
  toBinary,
  toJson,
} from "@bufbuild/protobuf";
import * as pb from "@/learningpb/addressbook_pb";
const book = ref<pb.AddressBook>(
  create(pb.AddressBookSchema, {
    people: [],
  })
);
const selectedPerson = ref<pb.Person[]>([
  create(pb.PersonSchema, {
    name: "",
    id: -1,
    email: "",
    phones: [],
  }),
]);

const personName = ref("");
const personEmail = ref("");
const personPhones = ref<pb.Person_PhoneNumber[]>([]);
const phoneNumber = ref("");
//const phoneType = ref<pb.PhoneType>(pb.PhoneType.UNSPECIFIED);
const phoneType = ref<pb.PhoneType>(pb.PhoneType.UNSPECIFIED);
const phoneInfo = ref<pb.Person_PhoneNumber[]>([]);

watch(selectedPerson, (newSelected, oldSelected) => {
  if (selectedPerson.value.length === 0) {
    return;
  }
  personName.value = selectedPerson.value[0].name;
  personEmail.value = selectedPerson.value[0].email;
  personPhones.value = selectedPerson.value[0].phones;

  phoneNumber.value = "";
  phoneType.value = pb.PhoneType.UNSPECIFIED;

  console.log("selected: ", selectedPerson.value[0]);
});
watch(phoneInfo, (newPhone, oldPhone) => {
  if (phoneInfo.value.length === 0) {
    return;
  }
  phoneNumber.value = phoneInfo.value[0].number;
  phoneType.value = phoneInfo.value[0].type;
});
let nc: nats_core.NatsConnection;
onMounted(async () => {
  nc = await nats_core.wsconnect({ servers: ["ws://127.0.0.1:9222"] });
  const subBookResponse = nc.subscribe("getBookResponse");
  (async () => {
    for await (const m of subBookResponse) {
      book.value = fromBinary(pb.AddressBookSchema, m.data);
      if (book.value.people.length === 0) {
        setInitBook();
      }
    }
  })();

  nc.publish(
    "getbookRequest",
    toBinary(pb.EmptySchema, create(pb.EmptySchema, {}))
  );
});

function savePhone() {
  personPhones.value.push(
    create(pb.Person_PhoneNumberSchema, {
      number: phoneNumber.value,
      type: phoneType.value,
    })
  );
  //if()
  phoneNumber.value = "";
  phoneType.value = pb.PhoneType.UNSPECIFIED;

  //console.log("savePhone");
}
function delPhone() {
  //console.log("delPhone");
  personPhones.value = personPhones.value.filter((item) => {
    return item.number !== phoneNumber.value;
  });
  phoneNumber.value = "";
}

async function savePerson() {
  console.log(
    "saving: ",
    personName.value,
    personEmail.value,
    personPhones.value
  );
  let i = 0;
  let id = 0;
  for (i = 0; i < book.value.people.length; i++) {
    if (book.value.people[i].id > id) {
      id = book.value.people[i].id;
    }
    if (book.value.people[i].name === personName.value) {
      book.value.people[i].email = personEmail.value;
      book.value.people[i].phones = personPhones.value;
      break;
    }
  }
  if (i === book.value.people.length) {
    book.value.people.push(
      create(pb.PersonSchema, {
        name: personName.value,
        id: id + 1,
        email: personEmail.value,
        phones: personPhones.value,
      })
    );
  }
  console.log("book: ", book.value);
  //return;
  nc.publish("savebookRequest", toBinary(pb.AddressBookSchema, book.value));
  resetInfo(true);
}
function delPerson() {
  book.value.people = book.value.people.filter((item) => {
    return item.name != personName.value;
  });
  //console.log("book: ", book.value);
  nc.publish("savebookRequest", toBinary(pb.AddressBookSchema, book.value));
  resetInfo(true);
}
function resetInfo(all: boolean) {
  personName.value = "";
  personEmail.value = "";
  phoneNumber.value = "";
  phoneType.value = pb.PhoneType.UNSPECIFIED;
  if (all) {
    personPhones.value = [];
    selectedPerson.value = [];
    phoneInfo.value = [];
  }
}
function setInitBook() {
  book.value.people.push(
    create(pb.PersonSchema, {
      name: "sunny",
      id: 1,
      email: "sunny@ever73.com",
    })
  );
  book.value.people.push(
    create(pb.PersonSchema, {
      name: "sam",
      id: 2,
      email: "sam@ever73.com",
      phones: [
        create(pb.Person_PhoneNumberSchema, {
          number: "13006505588",
          type: pb.PhoneType.WORK,
        }),
        create(pb.Person_PhoneNumberSchema, {
          number: "13001670032",
          type: pb.PhoneType.HOME,
        }),
      ],
    })
  );
  book.value.people.push(
    create(pb.PersonSchema, {
      name: "jf10",
      id: 3,
      email: "jf10@ever73.com",
    })
  );
}
</script>

<style scoped>
.outlined-row {
  border: 2px solid #1976d2; /* You can change the color and width as needed */
  border-radius: 4px; /* Optional: to give rounded corners */
  padding: 16px; /* Optional: to add some padding inside the row */
}
</style>
