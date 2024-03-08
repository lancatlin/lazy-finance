import { Transaction, Template } from "./types";

export function applyTemplate(
  tx: Transaction,
  template: Template
): Transaction {
  const accounts = tx.accounts.map((account, index) => ({
    name: account.name || template.accounts[index]?.name || "",
    amount: account.amount || template.accounts[index]?.amount,
    commodity: account.commodity || template.accounts[index]?.commodity || "",
  }));
  return {
    ...tx,
    accounts,
  };
}
