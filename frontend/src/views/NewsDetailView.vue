<template>
  <el-container>
    <el-main>
      <el-card v-if="article" class="article-detail">
        <div class="article-header">
          <h1 class="article-title">{{ article.Title }}</h1>
          <el-divider></el-divider>
        </div>
        <div class="article-content">
          <p>{{ article.Content }}</p>
        </div>
        <el-divider></el-divider>
        <div class="article-actions">
          <el-button type="primary" icon="el-icon-thumb" size="large" @click="likeArticle" class="like-button">点赞</el-button>
          <el-tag type="success" class="likes-count" effect="dark">点赞数: {{ likes }}</el-tag>
        </div>
      </el-card>
      <el-empty v-else class="no-data" description="您必须登录/注册才可以阅读文章"></el-empty>
    </el-main>
  </el-container>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import { useRoute } from "vue-router";
import axios from "../axios";
import type { Article, Like } from "../types/Article";

const article = ref<Article | null>(null);
const route = useRoute();
const likes = ref<number>(0);

const { id } = route.params;

const fetchArticle = async () => {
  try {
    const response = await axios.get<Article>(`/articles/${id}`);
    article.value = response.data;
  } catch (error) {
    console.error("Failed to load article:", error);
  }
};

const likeArticle = async () => {
  try {
    await axios.post<Like>(`/articles/${id}/like`);
    await fetchLike();
  } catch (error) {
    console.log("Error Liking article:", error);
  }
};

const fetchLike = async () => {
  try {
    const res = await axios.get<Like>(`/articles/${id}/like`);
    likes.value = res.data.likes;
  } catch (error) {
    console.log("Error fetching likes:", error);
  }
};

onMounted(fetchArticle);
onMounted(fetchLike);
</script>

<style scoped>
.article-detail {
  margin: 40px auto;
  padding: 30px;
  max-width: 800px;
  background-color: #ffffff;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.article-header {
  text-align: center;
  margin-bottom: 20px;
}

.article-title {
  font-size: 32px;
  font-weight: bold;
  color: #333;
}

.article-content {
  font-size: 18px;
  line-height: 1.8;
  color: #555;
  margin-bottom: 20px;
}

.article-actions {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 20px;
  margin-top: 20px;
}

.like-button {
  font-size: 16px;
  padding: 10px 20px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.likes-count {
  font-size: 16px;
  background-color: #eaf4e9;
  color: #2e7d32;
  display: flex;
  align-items: center;
  justify-content: center;
  height: 40px;
}

.no-data {
  text-align: center;
  font-size: 1.4em;
  color: #999;
  margin-top: 40px;
}
</style>
