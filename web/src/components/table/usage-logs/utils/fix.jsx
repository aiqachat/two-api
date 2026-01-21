// 修复描述"日志详情"
import { getCurrencyConfig } from '@/helpers/index.js';

export const fixDescriptionsLogDetails = (info, source) => {
  const details = info?.video_model_ratio_info;
  if (!details) {
    return source;
  }
  // 获取货币配置
  const { symbol } = getCurrencyConfig();
  const {
    price,
    group_ratio,
    duration,
    resolution,
    generate_audio_ratio,
    draft_ratio,
    service_tier_flex_ratio,
  } = details;
  let str = `分组倍率 ${group_ratio}; 视频每秒价格 ${symbol}${price}; 秒数 ${duration}; 分辨率 ${resolution};`;
  if (generate_audio_ratio) {
    str += ` 生成声音倍率: ${generate_audio_ratio};`;
  }
  if (draft_ratio) {
    str += ` 样片倍率: ${draft_ratio};`;
  }
  if (service_tier_flex_ratio) {
    str += ` 离线推理模式倍率: ${service_tier_flex_ratio};`;
  }
  return str;
};
// 修复描述"计费过程"
export const fixDescriptionsCalculateProcess = (info, source) => {
  const details = info?.video_model_ratio_info;
  if (!details) {
    return source;
  }
  // 获取货币配置
  const { symbol } = getCurrencyConfig();
  const {
    price,
    group_ratio,
    duration,
    resolution,
    generate_audio_ratio,
    draft_ratio,
    service_tier_flex_ratio,
    price_total,
  } = details;
  let str = `分组倍率：${group_ratio} * 分辨率(${resolution})价格：${symbol}${price} * 视频秒数：${duration}`;
  if (generate_audio_ratio) {
    str += ` * 生成声音倍率: ${generate_audio_ratio}`;
  }
  if (draft_ratio) {
    str += ` * 样片倍率: ${draft_ratio}`;
  }
  if (service_tier_flex_ratio) {
    str += ` * 离线推理模式倍率: ${service_tier_flex_ratio}`;
  }
  return `${str} = ${symbol}${price_total}`;
};
