<template>
  <el-container>
    <el-card class="exchange-form-card">
      <h2 class="form-title">货币兑换工具</h2>
      <p class="form-description">快速便捷地转换不同货币之间的金额。</p>

      <el-form :model="form" class="exchange-form">
        <el-form-item label="从货币" label-width="100px">
          <el-select v-model="form.fromCurrency" placeholder="选择货币" class="currency-select">
            <el-option v-for="currency in currencies" :key="currency" :label="currency" :value="currency" />
          </el-select>
        </el-form-item>
        <el-form-item label="到货币" label-width="100px">
          <el-select v-model="form.toCurrency" placeholder="选择货币" class="currency-select">
            <el-option v-for="currency in currencies" :key="currency" :label="currency" :value="currency" />
          </el-select>
        </el-form-item>
        <el-form-item label="金额" label-width="100px">
          <el-input v-model="form.amount" type="number" placeholder="输入金额" class="amount-input" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" size="large" @click="exchange" class="exchange-button">立即兑换</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card v-if="result !== null" class="result-card">
      <el-result icon="el-icon-s-promotion" title="兑换结果">
        <template #sub-title>
          <span class="result-text">{{ result }} {{ form.toCurrency }}</span>
        </template>
      </el-result>
    </el-card>

    <el-card v-if="historicalRates.length" class="history-section">
      <h3 class="history-title">最近兑换记录</h3>
      <el-table :data="historicalRates.slice(-5).reverse()" style="width: 100%">
        <el-table-column prop="fromAmount" label="兑换金额" width="120" />
        <el-table-column prop="fromCurrency" label="从货币" width="120" />
        <el-table-column prop="toAmount" label="兑换结果" width="120" />
        <el-table-column prop="toCurrency" label="到货币" width="120" />
      </el-table>
    </el-card>
  </el-container>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import axios from '../axios';

interface ExchangeRate {
  fromCurrency: string;
  toCurrency: string;
  rate: number;
}

interface ExchangeRecord {
  fromCurrency: string;
  toCurrency: string;
  fromAmount: number;
  toAmount: number;
}

const form = ref({
  fromCurrency: '',
  toCurrency: '',
  amount: 0,
});

const result = ref<number | null>(null);
const currencies = ref<string[]>([]);
const rates = ref<ExchangeRate[]>([]);
const historicalRates = ref<ExchangeRecord[]>([]);

const fetchCurrencies = async () => {
  try {
    const response = await axios.get<ExchangeRate[]>('/exchangeRates');
    rates.value = response.data;
    currencies.value = [...new Set(response.data.map((rate: ExchangeRate) => [rate.fromCurrency, rate.toCurrency]).flat())];
  } catch (error) {
    console.log('Failed to load currencies', error);
  }
};

const exchange = () => {
  const rate = rates.value.find(
    (rate) => rate.fromCurrency === form.value.fromCurrency && rate.toCurrency === form.value.toCurrency
  )?.rate;

  if (rate) {
    result.value = form.value.amount * rate;
    saveToHistory(form.value.fromCurrency, form.value.toCurrency, form.value.amount, result.value);
  } else {
    result.value = null;
  }
};

const saveToHistory = (fromCurrency: string, toCurrency: string, fromAmount: number, toAmount: number) => {
  historicalRates.value.push({ fromCurrency, toCurrency, fromAmount, toAmount });
};

onMounted(fetchCurrencies);
</script>

<style scoped>
.exchange-form-card {
  max-width: 600px;
  margin: 40px auto;
  padding: 30px;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  background-color: #ffffff;
}

.form-title {
  font-size: 28px;
  font-weight: bold;
  color: #333;
  text-align: center;
  margin-bottom: 10px;
}

.form-description {
  font-size: 16px;
  color: #666;
  text-align: center;
  margin-bottom: 30px;
}

.currency-select {
  width: 100%;
}

.amount-input {
  width: 100%;
}

.exchange-button {
  display: block;
  width: 100%;
  font-size: 18px;
}

.result-card {
  max-width: 600px;
  margin: 20px auto;
  padding: 20px;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  background-color: #ffffff;
}

.result-text {
  font-weight: bold;
  font-size: 20px;
  color: #2e7d32;
}

.history-section {
  max-width: 600px;
  margin: 40px auto;
  padding: 20px;
  border-radius: 12px;
  background-color: #ffffff;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.history-title {
  font-size: 22px;
  font-weight: bold;
  color: #333;
  margin-bottom: 20px;
}
</style>
