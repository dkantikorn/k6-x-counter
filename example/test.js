import { nextId, value, reset } from "k6/x/counter";

export function setup() {
  reset(); // Start counting from zero for each test run.
}

export default function () {
  for (let i = 0; i < 3; i++) {
    const id = nextId(6); // "000001", "000002", ...
    console.log(`VU ${__VU} got ID: ${id}`);
  }
}

export function teardown() {
  console.log(`Total IDs issued: ${value()}`);
}
