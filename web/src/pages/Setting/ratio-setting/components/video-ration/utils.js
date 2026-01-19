
export const getResolutionValue = (resolution) => {
  return +resolution.replace(/(\d+)[\s\S]*$/, '$1');
};