export const getVideoDownloadUrl = (info) => {
  if (!info) return '';
  const { data, code, content } = info;
  if (code === 10000) {
    return data?.video_url;
  }
  return content?.video_url;
};
