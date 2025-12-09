import { deerService, WsError } from '@helpers';

const getWsVideoRationPageList = async ({ pageSize, pageNumber }) => {
  try {
    const res = await deerService.getPageList('/api/ws/video-ratio/page', {
      pageSize,
      pageNumber,
    });
    WsError.checkApiResult(res);
    return res.data;
  } catch (e) {
    WsError.handleError(e);
    return [];
  }
};

export default {
  getWsVideoRationPageList,
};
