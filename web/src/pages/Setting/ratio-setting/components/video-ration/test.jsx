import { Modal } from '@douyinfe/semi-ui';
import { wsCreateModalHandle } from '@components';

export const TestModal = ({
                            modalProps,
                            onComplete,
                          }) => {
  return (
    <Modal
      {...modalProps}
      title='Basic Modal'
      onOk={() => {
        console.log('-----1');
        onComplete(true)
        modalProps.onCancel()
      }}
      // onCancel={() => {
      //   console.log('-----2');
      // }}
    >
      <p>Some contents...</p>
      <p>Some contents...</p>
      <p>Some contents...</p>
    </Modal>
  );
};

export const testModal = wsCreateModalHandle(TestModal);
