import { deerService, WsError } from '@helpers';

const getWsVideoRationDetails = async (id) => {
  const res = await deerService.post('/api/ws/video-ratio/details', {
    id,
  });
  WsError.checkApiResult(res);
  return res.data;
};

const createWsVideoRation = async ({ model_name, ...config }) => {
  try {
    const res = await deerService.post('/api/ws/video-ratio/create', {
      model_name,
      config,
    });
    WsError.checkApiResult(res);
  } catch (e) {
    WsError.handleError(e);
  }
};

const getResolutionOptionsList = async () => {
  try {
    const res = await deerService.post(
      '/api/ws/video-ratio/resolutionList',
      {},
    );
    WsError.checkApiResult(res);
    return res.data.items;
  } catch (e) {
    WsError.handleError(e);
    return [];
  }
};

const getModelOptionsList = async () => {
  try {
    const res = await deerService.getPageList('/api/models/', {
      page_size: 10000,
    });
    WsError.checkApiResult(res);
    return res.data.items.map((model) => {
      return {
        label: model.model_name,
        value: model.model_name,
      };
    });
  } catch (e) {
    WsError.handleError(e);
    return [];
  }
};

export default {
  getWsVideoRationDetails,
  getModelOptionsList,
  getResolutionOptionsList,
  createWsVideoRation,
};
