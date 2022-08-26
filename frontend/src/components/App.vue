<template>
  <div>
    <h1 class="text-5xl">Go Web Application</h1>
    <login-component></login-component>
  </div>
</template>

<script>
import LoginComponent from './Login.vue';
export default {
  components: {
    LoginComponent
  },
  methods: {
    login() {

    }
  }
}
</script>