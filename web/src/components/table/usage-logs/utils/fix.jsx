// 修复描述"日志详情"
export const fixDescriptionsLogDetails = (info, source) => {
  const details = info?.video_model_ratio_info;
  if (!details) {
    return source;
  }
  const { price, group_ratio, duration, resolution } = details;
  return `分组倍率 ${group_ratio}，视频每秒价格 $${price}，秒数 ${duration}，分辨率 ${resolution}`;
};
// 修复描述"计费过程"
export const fixDescriptionsCalculateProcess = (info, source) => {
  const details = info?.video_model_ratio_info;
  if (!details) {
    return source;
  }
  const { price, group_ratio, duration, resolution, price_total } = details;
  return `分组倍率：${group_ratio} * 分辨率(${resolution})价格：$${price} * 视频秒数：${duration} = $${price_total}`;
};