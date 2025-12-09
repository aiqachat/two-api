import { deerService, WsError } from '@helpers';

const createWsVideoRation = async ({ modeName, resolution, price }) => {
  try {
    const res = await deerService.post('/api/ws/video-ratio/create', {
      page_size: 10000,
      p: 1,
      mode_name: modeName,
      config: {
        [resolution]: price
      }
    });
    WsError.checkApiResult(res);
  } catch (e) {
    WsError.handleError(e);
  }
};

const getFullModelList = async () => {
  try {
    const res = await deerService.getPageList('/api/models/', {
      page_size: 10000,
    });
    WsError.checkApiResult(res);
    return res.data.items;
  } catch (e) {
    WsError.handleError(e);
    return [];
  }
};

export default {
  getFullModelList,
  createWsVideoRation,
};
