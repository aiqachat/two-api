import { Form, Modal } from '@douyinfe/semi-ui';
import { wsCreateModalHandle } from '@components';
import { useEffect, useRef } from 'react';
import { WsError } from '@helpers';
import service from './service';
import { RESOLUTION_LIST } from '../../const';
import { useAsync } from 'react-use';

export const EditModal = ({ modalProps, onComplete, edit = true, id }) => {
  const formRef = useRef(null);
  const fullModels = useAsync(service.getFullModelList);
  const modelOptions = (fullModels.value || []).map((model) => {
    return {
      label: model.model_name,
      value: model.model_name,
    };
  });
  useEffect(() => {
    formRef.current?.formApi.setValues({
      name: '视频比率',
      resolution: RESOLUTION_LIST[0].value,
    });
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
          field='modeName'
          rules={[{ required: true }]}
          loading={fullModels.loading}
          optionList={modelOptions}
          placeholder='请选择模型'
          style={{width: '100%'}}
        />
        <Form.Select
          label='分辨率'
          field='resolution'
          rules={[{ required: true }]}
          optionList={RESOLUTION_LIST}
          placeholder='请选择分辨率'
          style={{width: 200}}
        />
        <Form.InputNumber
          label='每秒价格'
          field='price'
          rules={[{ required: true }]}
          precision={2}
          step={1}
          min={0}
          max={99999999999}
          placeholder='请输入每秒价格'
        />
      </Form>
    </Modal>
  );
};

export const editModal = wsCreateModalHandle(EditModal);
