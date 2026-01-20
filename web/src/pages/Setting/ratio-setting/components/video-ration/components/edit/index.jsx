import { Form, Modal } from '@douyinfe/semi-ui';
import { wsCreateModalHandle } from '@components';
import { useEffect, useRef } from 'react';
import _ from 'lodash';
import { WsError } from '@helpers';
import service from './service';
import { useAsync } from 'react-use';

export const EditModal = ({ modalProps, onComplete, edit = true, id }) => {
  const formRef = useRef(null);
  const modelRes = useAsync(service.getModelOptionsList);
  const initConfigRes = useAsync(service.getWsVideoRatioInitConfig);

  const initConfigList = initConfigRes.value || [];

  const loadDetails = async () => {
    try {
      const res = await service.getWsVideoRationDetails(id);
      const config = {};
      res.config.forEach(({ name, value }) => {
        config[name] = value;
      });
      formRef.current?.formApi.setValues({
        model_name: res.model_name,
        config,
      });
    } catch (e) {
      WsError.handleError(e);
    }
  };

  useEffect(() => {
    if (initConfigList.length === 0) return;
    if (!modelRes.value) return;
    formRef.current?.formApi.setValues({
      config: _.fromPairs(
        initConfigList.map(({ name, value }) => [name, value]),
      ),
    });
    if (!edit) return;
    loadDetails().then();
  }, [edit, initConfigList, modelRes.value]);

  return (
    <Modal
      {...modalProps}
      title='编辑视频比率'
      onOk={async () => {
        try {
          const { model_name, config } = await formRef.current?.formApi.validate();
          const values = {
            model_name,
            config: initConfigList.map((item) => ({
              ...item,
              value: config[item.name],
            })),
          };
          if (edit) {
            console.log(values);
            await service.editWsVideoRation({
              id,
              ...values,
            });
          } else {
            await service.createWsVideoRation(values);
          }
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
          disabled={edit}
          rules={[{ required: true }]}
          loading={modelRes.loading}
          optionList={(modelRes.value || []).map((model) => {
            return {
              label: model.model_name,
              value: model.model_name,
            };
          })}
          placeholder='请选择模型'
          style={{ width: '100%' }}
        />
        {initConfigList.map(({ label, name }) => {
          return (
            <Form.InputNumber
              label={label}
              field={`config.${name}`}
              rules={[{ required: true }]}
              precision={3}
              step={1}
              min={0}
              max={99999999999}
              placeholder={`请输入${label}`}
            />
          );
        })}
      </Form>
    </Modal>
  );
};

export const editModal = wsCreateModalHandle(EditModal);
