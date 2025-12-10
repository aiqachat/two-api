export * from './fix';

export const renderVideoRatioInfo = (info) => {
  return `价格: $${info.price} * 分组倍率: ${info.user_group_ratio} * 秒数: ${info.duration}`;
};
