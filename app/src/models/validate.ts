import { Transaction, Account } from "./types";

export default function validate(tx: Transaction): Transaction {
  let nullCount = 0;
  let totalAmount = 0;
  const { template } = tx;
  const accounts = tx.accounts.map((account, index) => {
    const output: Account = {
      name: account.name || template.accounts[index]?.name || "",
      amount: account.amount || template.accounts[index]?.amount,
      commodity: account.commodity || template.accounts[index]?.commodity || "",
    };
    if (output.amount === null) nullCount++;
    else totalAmount += output.amount as number; // Safe cast since we check for null above
    if (!output.name) {
      throw `Account ${index + 1} cannot be empty.`;
    }
    return output;
  });
  console.log(accounts);

  if (nullCount > 1) {
    throw "Only one account can have an empty amount.";
  } else if (nullCount === 0 && totalAmount !== 0) {
    throw "The total amount of every account should be 0.";
  }
  return {
    ...tx,
    accounts,
  };
}
