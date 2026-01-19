import { getCurrencyConfig } from '@/helpers/index.js';

export * from './fix';

export const renderVideoRatioInfo = (info) => {
  // 获取货币配置
  const { symbol } = getCurrencyConfig();
  return `价格: ${symbol}${info.price} * 分组倍率: ${info.group_ratio} * 秒数: ${info.duration}`;
};