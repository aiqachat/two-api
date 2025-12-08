import { useEffect, useState } from 'react'
import { render, unmount } from 'rc-util/es/React/render'

const ModalHandleContainer = ({
  Comp,
  data,
  onComplete,
  onUnmount,
  r
}) => {
  const [ visible, setVisible ] = useState(true)
  useEffect(() => {
    setVisible(true)
  }, [ r ])
  return (
    <>
      <Comp
        modalProps={{
          destroyOnHidden: true,
          onCancel: () => {
            setVisible(false)
            onUnmount()
          },
          visible,
        }}
        onComplete={onComplete || (() => undefined)}
        {...data}
      />
    </>
  )
}

/**
 * 创建弹窗控制器
 */
export const wsCreateModalHandle = (Comp) => {
  const container = document.createDocumentFragment()
  return {
    open: (data, onComplete) => {
      render((
          <ModalHandleContainer
            Comp={Comp}
            data={data}
            onComplete={onComplete}
            onUnmount={() => {
              setTimeout(() => {
                unmount(container).then()
              }, 300)
            }}
            r={Math.random()}
          />
      ), container)
    }
  }
}
