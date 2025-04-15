import { Slot } from "expo-router";
import { StatusBar } from "expo-status-bar";

export default function Root() {
  return (
    <>
      <StatusBar style="dark" />
      {/* Authentification provider */}
      <Slot />
      {/* Authentification provider */}
    </>
  );
}
