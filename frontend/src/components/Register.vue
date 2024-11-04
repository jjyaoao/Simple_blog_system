<template>  
  <div class="auth-container">  
    <el-card class="auth-form-card">
      <h2 class="auth-title">用户注册</h2>
      <el-form :model="form" class="auth-form" @submit.prevent="register">  
        <el-form-item label="用户名" label-width="100px">  
          <el-input v-model="form.username" placeholder="请输入用户名" prefix-icon="el-icon-user" />  
        </el-form-item>  
        <el-form-item label="密码" label-width="100px">  
          <el-input v-model="form.password" type="password" placeholder="请输入密码" prefix-icon="el-icon-lock" />  
        </el-form-item>  
        <el-form-item>  
          <el-button type="primary" native-type="submit" size="large" class="register-button">注册</el-button>  
        </el-form-item>  
      </el-form>  
    </el-card>
  </div>  
</template>  
  
<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { useAuthStore } from '../store/auth';
import { ElMessage } from 'element-plus';

const form = ref({
  username: '',
  password: '',
});

const authStore = useAuthStore();
const router = useRouter();

const register = async () => {
  try {
    await authStore.register(form.value.username, form.value.password);
    router.push({ name: 'News' });
  } catch {
    ElMessage.error('注册失败，请重试。');
  }
};
</script>
  
<style scoped>
.auth-container {  
  display: flex;  
  justify-content: center;  
  align-items: center;  
  height: 100vh; 
  background: linear-gradient(135deg, #f0f2f5, #d9e7ff); 
  padding: 20px;  
  box-sizing: border-box; 
}  
  
.auth-form-card {
  width: 100%;  
  max-width: 400px; 
  padding: 30px;  
  background-color: #ffffff;  
  border-radius: 12px;  
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);  
}

.auth-title {
  text-align: center;
  font-size: 24px;
  font-weight: bold;
  color: #333;
  margin-bottom: 20px;
}

.auth-form {
  margin-top: 20px;
}

.register-button {
  width: 100%;
  font-size: 18px;
  padding: 10px;
}
</style>
