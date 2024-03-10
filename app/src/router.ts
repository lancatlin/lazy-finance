import { createRouter, createWebHistory } from "vue-router";
import Home from "./views/HomePage.vue";
import Edit from "./views/EditPage.vue";
import Transactions from "./views/TransactionPage.vue";
import Balances from "./views/BalancePage.vue";
import SignInPage from "./views/SignInPage.vue";
import SignUpPage from "./views/SignUpPage.vue";
import { isSignedIn } from "./utils/api";

const routes = [
  {
    path: "/",
    redirect: "/dashboard",
  },
  {
    path: "/dashboard",
    name: "Home",
    component: Home,
    meta: { requiresAuth: true },
  },
  {
    path: "/edit",
    name: "EditHome",
    component: Edit,
    meta: { requiresAuth: true },
  },
  {
    path: "/edit/:path",
    name: "EditWithPath",
    component: Edit,
    meta: { requiresAuth: true },
  },
  {
    path: "/transactions",
    name: "Transactions",
    component: Transactions,
    meta: { requiresAuth: true },
  },
  {
    path: "/balances",
    name: "Balances",
    component: Balances,
    meta: { requiresAuth: true },
  },
  {
    path: "/signin",
    name: "SignIn",
    component: SignInPage,
    meta: { requiresGuest: true },
  },
  {
    path: "/signup",
    name: "SignUp",
    component: SignUpPage,
    meta: { requiresGuest: true },
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach(async (to, _, next) => {
  const signedIn = await isSignedIn();
  // Check if the route requires authentication
  if (to.meta.requiresAuth && !signedIn) {
    // Check if the user isn't authenticated
    // Redirect to the login page
    next({
      name: "SignIn",
      query: { redirect: to.fullPath }, // Optionally pass a redirect parameter
    });
  } else if (to.meta.requiresGuest && signedIn) {
    next({
      name: "Home",
    });
  } else {
    next();
  }
});

export default router;
