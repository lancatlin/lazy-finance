import { reactive } from "vue";
import { Session } from "./types";

const session = reactive<Session>({
  signedIn: false,
});

export function emitSignIn() {
  session.signedIn = true;
}

export function emitSignOut() {
  session.signedIn = false;
}

export default session;
