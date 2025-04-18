import { Platform } from "react-native";
import axios, { AxiosInstance } from "axios";

const url =
  Platform.OS === "android" ? "http://10.0.2.2:3000" : "http://127.0.0.1:3000";

  const Api: AxiosInstance = axios.create({
    baseURL: url + "api",

  })