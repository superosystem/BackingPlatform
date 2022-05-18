import { createWebHistory, createRouter } from "vue-router";

const routes = [
    {
        path: "/",
        alias: "/tutorials",
        name: "tutorials",
        component: () => import("./components/TutorialsList.vue")
    },
    {
        path: "/tutorials/:id",
        name: "tutorials-details",
        component: () => import("./components/Tutorial.vue")
    },
    {
        path: "/add",
        name: "add",
        component: () => import("./components/AddTutorial.vue")
    },
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});
export default router;
