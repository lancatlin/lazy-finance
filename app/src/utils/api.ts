import axios from "axios";
import { Balance, Template, Transaction } from "../models/types";

const api = axios.create({
  baseURL: "/api",
  timeout: 2000,
});

export async function getTemplates(): Promise<Template[]> {
  const response = await api.get<Template[]>("/templates");
  return response.data;
}

export async function newTx(data: Transaction): Promise<Transaction> {
  try {
    const response = await api.post<Transaction>("/txs", data);
    return response.data;
  } catch (error) {
    if (axios.isAxiosError(error) && error.response?.status === 400) {
      throw new Error(
        error.response.data.message ||
          "An error occurred while trying to create a new transaction."
      );
    }
    throw error; // Rethrow error if it's not the specific 400 error we're checking for or if it's a different kind of error
  }
}

export async function getTxs(): Promise<Transaction[]> {
  const response = await api.get<Transaction[]>("/txs");
  const txs = response.data;
  return txs.map((tx) => ({
    ...tx,
    date: new Date(tx.date),
  }));
}

export async function getBalances(): Promise<Balance[]> {
  const response = await api.get<Balance[]>("/balances");
  return response.data;
}
