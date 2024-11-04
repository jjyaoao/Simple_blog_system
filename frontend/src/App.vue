<template>
  <el-container>
    <el-aside width="240px" class="nav-aside">
      <el-menu :default-active="activeIndex" class="navigation-menu" mode="vertical" @select="handleSelect">
        <el-menu-item index="home"><i class="el-icon-house"></i> 首页</el-menu-item>
        <el-menu-item index="currencyExchange"><i class="el-icon-money"></i> 兑换货币</el-menu-item>
        <el-menu-item index="news"><i class="el-icon-document"></i> 查看新闻</el-menu-item>
        <el-menu-item index="login" v-if="!authStore.isAuthenticated"><i class="el-icon-user"></i> 登录</el-menu-item>
        <el-menu-item index="register" v-if="!authStore.isAuthenticated"><i class="el-icon-edit"></i> 注册</el-menu-item>
        <el-menu-item index="logout" v-if="authStore.isAuthenticated"><i class="el-icon-switch-button"></i> 退出</el-menu-item>
      </el-menu>
    </el-aside>
    <el-container>
      <el-header class="header">
        <div class="header-title">轻博客</div>
      </el-header>
      <el-main>
        <router-view></router-view>
      </el-main>
    </el-container>
  </el-container>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import { useAuthStore } from './store/auth';

const router = useRouter();
const route = useRoute();
const authStore = useAuthStore();
const activeIndex = ref(route.name?.toString() || 'home');

watch(route, (newRoute) => {
  activeIndex.value = newRoute.name?.toString() || 'home';
});

const handleSelect = (key: string) => {
  if (key === 'logout') {
    authStore.logout();
    router.push({ name: 'Home' });
  } else {
    router.push({ name: key.charAt(0).toUpperCase() + key.slice(1) });
  }
};
</script>

<style scoped>
.navigation-menu {
  background-color: #ffffff;
  color: #333;
  border-right: 1px solid #e0e0e0;
  height: 100%;
  padding-top: 20px;
}

.navigation-menu .el-menu-item {
  font-size: 16px;
  padding: 15px 20px;
  transition: background-color 0.3s, color 0.3s;
  border-radius: 8px;
  margin: 5px;
}

.navigation-menu .el-menu-item:hover {
  background-color: #f0f0f0;
}

.navigation-menu .el-menu-item.is-active {
  background-color: #409EFF;
  color: #ffffff;
}

.navigation-menu .el-icon-house,
.navigation-menu .el-icon-money,
.navigation-menu .el-icon-document,
.navigation-menu .el-icon-user,
.navigation-menu .el-icon-edit,
.navigation-menu .el-icon-switch-button {
  margin-right: 12px;
}

.nav-aside {
  background-color: #f9f9f9;
  border-right: 1px solid #e0e0e0;
}

.header {
  display: flex;
  align-items: center;
  padding-left: 20px;
  background-color: #ffffff;
  border-bottom: 1px solid #e0e0e0;
  height: 60px;
}

.header-title {
  font-size: 20px;
  font-weight: bold;
  color: #333;
}
</style>
