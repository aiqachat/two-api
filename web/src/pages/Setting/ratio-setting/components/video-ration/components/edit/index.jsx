import { Form, Modal } from '@douyinfe/semi-ui';
import { wsCreateModalHandle } from '@components';
import { useEffect, useRef } from 'react';
import { WsError } from '@helpers';
import service from './service';
import { useAsync } from 'react-use';

export const EditModal = ({ modalProps, onComplete, edit = true, id }) => {
  const formRef = useRef(null);
  const modelOptions = useAsync(service.getModelOptionsList);
  const resolutionItems = useAsync(service.getResolutionOptionsList);
  const loadDetails = async () => {
    try {
      const res = await service.getWsVideoRationDetails(id);
      console.log(res)
      formRef.current?.formApi.setValues({
        model_name: res.model_name,
        // resolution: RESOLUTION_LIST[0].value,
      });
    } catch (e) {
      WsError.handleError(e);
    }
  };
  useEffect(() => {
    if (!edit) return;
    loadDetails().then();
  }, [edit]);
  return (
    <Modal
      {...modalProps}
      title='编辑视频比率'
      onOk={async () => {
        try {
          const values = await formRef.current?.formApi.validate();
          await service.createWsVideoRation(values);
          onComplete(true);
          modalProps.onCancel();
        } catch (e) {
          WsError.handleError(e);
        }
      }}
    >
      <Form ref={formRef}>
        <Form.Select
          label='模型'
          field='model_name'
          rules={[{ required: true }]}
          loading={modelOptions.loading}
          optionList={modelOptions.value}
          placeholder='请选择模型'
          style={{ width: '100%' }}
        />
        {(resolutionItems.value || []).map((item) => {
          return (
            <Form.InputNumber
              label={`分辨率${item.name}每秒价格`}
              field={item.key}
              rules={[{ required: true }]}
              precision={2}
              step={1}
              min={0}
              max={99999999999}
              placeholder='请输入价格'
            />
          );
        })}
      </Form>
    </Modal>
  );
};

export const editModal = wsCreateModalHandle(EditModal);
