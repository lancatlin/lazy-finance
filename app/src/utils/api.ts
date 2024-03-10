import axios from "axios";
import {
  Balance,
  Template,
  Transaction,
  Query,
  File,
  Session,
  LoginRequest,
} from "../models/types";
import { emitSignIn, emitSignOut } from "../models/sessionStore";

const api = axios.create({
  baseURL: "/api",
  timeout: 2000,
});

export async function getTemplates(): Promise<Template[]> {
  const response = await api.get<Template[]>("/templates");
  return response.data;
}

export async function newTx(
  data: Transaction,
  options: { save: boolean }
): Promise<Transaction> {
  try {
    const response = await api.post<Transaction>("/txs", data, {
      params: options,
    });
    return response.data;
  } catch (error) {
    throw handleHttpError(error);
  }
}

function handleHttpError(error: unknown): any {
  if (axios.isAxiosError(error) && error.response?.status === 400) {
    return new Error(
      error.response.data.message ||
        "An error occurred while trying to create a new transaction."
    );
  }
  return error;
}

export async function getTxs(query: Query): Promise<Transaction[]> {
  const response = await api.get<Transaction[]>("/txs", { params: query });
  const txs = response.data;
  return txs.map((tx) => ({
    ...tx,
    date: new Date(tx.date),
  }));
}

export async function getBalances(query: Query): Promise<Balance[]> {
  const response = await api.get<Balance[]>("/balances", { params: query });
  return response.data;
}

export async function getFileList(): Promise<File[]> {
  const response = await api.get<File[]>("/files");
  return response.data;
}

export async function getFileContent(path: string): Promise<string> {
  const response = await api.get<string>(`/files/${path}`, {
    responseType: "text",
  });
  return response.data;
}

export async function saveFile(path: string, content: string): Promise<void> {
  await api.post(`/files/${path}`, { data: content });
}

export async function signIn(payload: LoginRequest): Promise<void> {
  try {
    await api.post("/signin", payload);
    emitSignIn();
  } catch (err) {
    throw handleAuthError(err);
  }
}

export async function signUp(payload: LoginRequest): Promise<void> {
  try {
    await api.post("/signup", payload);
    emitSignIn();
  } catch (err) {
    throw handleAuthError(err);
  }
}

function handleAuthError(err: unknown): any {
  if (axios.isAxiosError(err) && err.response?.status === 401) {
    return new Error("Invalid email or password");
  }
  return err;
}

export async function logout(): Promise<void> {
  await api.post("/logout");
  emitSignOut();
}

export async function isSignedIn(): Promise<boolean> {
  const response = await api.get<Session>("/status");
  if (response.data.signedIn) {
    emitSignIn();
  } else {
    emitSignOut();
  }
  return response.data.signedIn;
}
