<template>
  <el-container class="home-container">
    <div class="content-wrapper">
      <h1 class="title">欢迎来到jjy的轻博客</h1>
      <p class="description">探索、记录、分享，请从上方选择功能开始您的旅程。</p>
    </div>
  </el-container>
</template>

<script setup lang="ts">
</script>

<style scoped>
.home-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background: linear-gradient(135deg, #f9f9f9, #e0e0e0);
  padding: 20px;
  box-sizing: border-box;
}

.content-wrapper {
  text-align: center;
  max-width: 600px;
  background: white;
  padding: 40px;
  border-radius: 12px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

.title {
  color: #2c3e50;
  font-size: 32px;
  font-weight: 600;
  margin-bottom: 15px;
}

.description {
  color: #7f8c8d;
  font-size: 16px;
  line-height: 1.6;
  margin-bottom: 30px;
}
</style>
