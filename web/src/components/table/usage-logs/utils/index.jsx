import { getCurrencyConfig } from '@/helpers/index.js';

export * from './fix';

export const renderVideoRatioInfo = (info) => {
  // 获取货币配置
  const { symbol } = getCurrencyConfig();
  const {
    price,
    group_ratio,
    duration,
    generate_audio_ratio,
    draft_ratio,
    service_tier_flex_ratio,
  } = info;
  let str = `价格: ${symbol}${price} * 分组倍率: ${group_ratio} * 秒数: ${duration}`;
  if (generate_audio_ratio) {
    str += ` * 生成声音倍率: ${generate_audio_ratio}`;
  }
  if (draft_ratio) {
    str += ` * 样片倍率: ${draft_ratio}`;
  }
  if (service_tier_flex_ratio) {
    str += ` * 离线推理模式倍率: ${service_tier_flex_ratio}`;
  }
  return str;
};
