import { Form, Modal } from '@douyinfe/semi-ui';
import { wsCreateModalHandle } from '@components';
import {useCallback, useEffect, useRef, useState } from 'react';
import _ from 'lodash';
import { WsError } from '@helpers';
import service from './service';
import { useAsync } from 'react-use';
import { getResolutionValue } from '../../utils';

export const EditModal = ({ modalProps, onComplete, edit = true, id }) => {
  const formRef = useRef(null);
  const modelRes = useAsync(service.getModelOptionsList);
  const resolutionRes = useAsync(service.getResolutionOptionsList);
  const [resolutionItems, setResolutionItems] = useState([]);

  const loadResolutionItems = useCallback((modelName) => {
    const model = (modelRes.value || []).find(
        (model) => model.model_name === modelName,
    );
    const resolutionArr = (resolutionRes.value || []).sort((a, b) => {
      return getResolutionValue(a.key) - getResolutionValue(b.key);
    });
    if (model && _.get(model, 'bound_channels[0].type') === 54) {
      setResolutionItems([
        ...resolutionArr.map(({ key, name }) => {
          return {
            name: `${name}(无声)`,
            key,
          };
        }),
        ...resolutionArr.map(({ key, name }) => {
          return {
            name: `${name}(有声)`,
            key: `${key}_audio`,
          };
        }),
      ]);
      return
    }
    setResolutionItems(resolutionArr.map(({ key, name }) => {
      return {
        name: `${name}(无声)`,
        key,
      };
    }));
  }, [resolutionRes.value, modelRes.value, setResolutionItems]);

  const loadDetails = async () => {
    try {
      const res = await service.getWsVideoRationDetails(id);
      loadResolutionItems(res.model_name);
      formRef.current?.formApi.setValues({
        model_name: res.model_name,
        config: res.config,
      });
    } catch (e) {
      WsError.handleError(e);
    }
  };

  useEffect(() => {
    loadResolutionItems('-1');
    if (!edit) return;
    if (!resolutionRes.value) return;
    if (!modelRes.value) return;
    loadDetails().then();
  }, [edit, resolutionRes.value, modelRes.value]);

  return (
    <Modal
      {...modalProps}
      title='编辑视频比率'
      onOk={async () => {
        try {
          const values = await formRef.current?.formApi.validate();
          if (edit) {
            await service.editWsVideoRation({ ...values, id });
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
          onChange={(val) => {
            loadResolutionItems(val);
          }}
        />
        {resolutionItems.map((item) => {
          return (
            <Form.InputNumber
              label={`分辨率${item.name}每秒价格`}
              field={`config.${item.key}`}
              rules={[{ required: true }]}
              precision={3}
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
