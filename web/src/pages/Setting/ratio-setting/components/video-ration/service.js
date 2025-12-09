import { deerService, WsError } from '@helpers';

const getWsVideoRationPageList = async ({ modeName, resolution, price }) => {
  try {
    const res = await deerService.getPageList('/api/ws/video-ratio/page', {
      page_size: 10000,
    });
    WsError.checkApiResult(res);
    return res.data
  } catch (e) {
    WsError.handleError(e);
    return [];
  }
};

export default {
  getWsVideoRationPageList,
};
