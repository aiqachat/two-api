export const getVideoDownloadUrl = (info) => {
  if (!info) return '';
  const { data, code, content, output } = info;
  if (code === 10000) {
    return data?.video_url;
  }
  if (output?.video_url) {
    return output?.video_url;
  }
  return content?.video_url;
};
