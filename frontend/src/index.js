import  { createApp }  from 'vue';
import { createStore } from 'vuex';
import App from './components/App.vue';
import "./index.css";

const store = createStore({
    state() {
        return {
            isAuthenticated: false,
            jwtKey: ""
        }
    },
    mutations: {
        setUserLoggedIn(state) {
            state.isAuthenticated = true;
        },
        setUserLoggedOut(state) {
            state.isAuthenticated = false;
        },
        setJWTToken(state, token) {
            state.jwtToken = token;
        }
    }
});

const app = createApp(App);
app.use(store);
app.mount("#app");
