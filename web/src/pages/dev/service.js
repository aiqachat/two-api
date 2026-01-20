import { deerService, WsError } from '@helpers';

const fixVideoRatioConfig = async () => {
  const res = await deerService.post('/api/ws/video-ratio/fix', {});
  WsError.checkApiResult(res);
};

export default {
  fixVideoRatioConfig,
};
