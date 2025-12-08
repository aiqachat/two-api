import { deerService, WsError } from '@helpers';

const createWsVideoRation = async ({ modeName, resolution, price }) => {
  try {
    const res = await deerService.post('/api/ws/video-ratio/create', {
      modeName,
      resolution,
      price,
    });
    WsError.checkApiResult(res)
  } catch (e) {
    WsError.handleError(e)
  }
};

export default {
  createWsVideoRation
}
