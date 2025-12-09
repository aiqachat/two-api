import { deerService, WsError } from '@helpers';

const getWsVideoRationPageList = async ({
  pageSize,
  pageNumber,
  model_name,
}) => {
  try {
    const res = await deerService.getPageList('/api/ws/video-ratio/page', {
      pageSize,
      pageNumber,
      model_name,
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
