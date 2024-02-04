export interface Account {
  name: string;
  amount: number | null;
  commodity: string;
}

export interface Template {
  accounts: Account[];
}

export interface Transaction {
  name: string;
  accounts: Account[];
  template: Template;
}
