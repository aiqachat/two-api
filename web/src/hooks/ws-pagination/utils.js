/**
 * 格式化范围选择器数据
 * @param value
 */
export const formatRangeValue = (value) => {
  if (!value?.length) {
    return value
  }
  if (!value[0]?.toDate) {
    return value
  }
  return value?.map((item) => {
    return item?.toDate().valueOf()
  })
}

/**
 * 格式化form数据
 * @param values
 */
export const formatValues = (values) => {
  const res = { ...values }
  Object.keys(values).forEach((key) => {
    res[key] = formatRangeValue(res[key])
  })
  return res
}
