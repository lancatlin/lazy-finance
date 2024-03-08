export interface Account {
  name: string;
  amount: number | null;
  commodity: string;
}

export interface Template {
  name: string;
  accounts: Account[];
}

export interface Transaction {
  name: string;
  accounts: Account[];
  template: Template;
}

export const defaultTemplate: Template = {
  name: "default",
  accounts: [
    {
      name: "",
      amount: 0,
      commodity: "",
    },
    {
      name: "",
      amount: 0,
      commodity: "",
    },
  ],
};
