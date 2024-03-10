import { createRouter, createWebHistory } from "vue-router";
import Home from "./views/HomePage.vue";
import Edit from "./views/EditPage.vue";
import Transactions from "./views/TransactionPage.vue";
import Balances from "./views/BalancePage.vue";

const routes = [
  {
    path: "/",
    name: "Home",
    component: Home,
  },
  {
    path: "/edit",
    name: "EditHome",
    component: Edit,
  },
  {
    path: "/edit/:path",
    name: "EditWithPath",
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
