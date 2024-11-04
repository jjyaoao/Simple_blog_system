<template>
  <el-container>
    <el-main>
      <div v-if="articles && articles.length" class="articles-container">
        <el-card v-for="article in articles" :key="article.ID" class="article-card" shadow="hover">
          <h2 class="article-title">{{ article.Title }}</h2>
          <p class="article-preview">{{ article.Preview }}</p>
          <el-button type="text" @click="viewDetail(article.ID)" class="read-more-button">阅读更多</el-button>
        </el-card>
      </div>
      <el-empty v-else class="no-data" description="您必须登录/注册才可以查看文章"></el-empty>
    </el-main>
  </el-container>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { ElMessage } from 'element-plus';
import axios from '../axios';
import { useAuthStore } from '../store/auth';
import type { Article } from "../types/Article";

const articles = ref<Article[]>([]);
const router = useRouter();
const authStore = useAuthStore();

const fetchArticles = async () => {
  try {
    const response = await axios.get<Article[]>('/articles');
    articles.value = response.data;
  } catch (error) {
    console.error('Failed to load articles:', error);
  }
};

const viewDetail = (id: string) => {
  if (!authStore.isAuthenticated) {
    ElMessage.error('请先登录后再查看');
    return;
  }
  router.push({ name: 'NewsDetail', params: { id } });
};

onMounted(fetchArticles);
</script>

<style scoped>
.articles-container {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 20px;
  padding: 20px;
}

.article-card {
  transition: transform 0.3s;
  cursor: pointer;
}

.article-card:hover {
  transform: translateY(-10px);
}

.article-title {
  font-size: 24px;
  font-weight: 600;
  margin-bottom: 10px;
  color: #333;
}

.article-preview {
  font-size: 16px;
  line-height: 1.6;
  color: #666;
  margin-bottom: 15px;
}

.read-more-button {
  font-size: 14px;
  color: #409EFF;
  padding: 0;
}

.no-data {
  text-align: center;
  font-size: 1.4em;
  color: #999;
  margin-top: 40px;
}
</style>
