import { createRouter, createWebHistory } from "vue-router";
import Home from "./views/Home.vue";
import Edit from "./views/Edit.vue";
import Transactions from "./views/Transactions.vue";
import Balances from "./views/Balances.vue";

const routes = [
  {
    path: "/",
    name: "Home",
    component: Home,
  },
  {
    path: "/edit",
    name: "Edit",
    component: Edit,
  },
  {
    path: "/transactions",
    name: "Transactions",
    component: Transactions,
  },
  {
    path: "/balances",
    name: "Balances",
    component: Balances,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
