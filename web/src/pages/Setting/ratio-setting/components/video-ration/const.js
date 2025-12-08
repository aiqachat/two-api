
export const RESOLUTION_MAP = {
  '1080p': '1080p',
  '720p': '720p',
  '480p': '480p',
};

export const RESOLUTION_LIST = Object.entries(RESOLUTION_MAP).map(([value, label]) => ({
  label,
  value,
}));
