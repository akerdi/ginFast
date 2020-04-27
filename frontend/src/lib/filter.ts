import Vue from 'vue'
import moment from 'moment'

Vue.filter('date', (value, format = 'YYYY-MM-DD HH:mm') => {
  if (!value) return ''
  return moment(value).format(format)
})

Vue.filter('size', (size: number, flag) => {
  let result = ''
  if (isNaN(Number(size))) return size
  if (!size) {
    result = '0 B'
  } else if (size < 1024) {
    result = parseInt(size * 100 + '') / 100 + ' B'
  } else if (size < (1024 * 1024)) {
    result = parseInt((size / 1024) * 100 + '') / 100 + ' KB'
  } else if (size < (1024 * 1024 * 1024)) {
    result = parseInt((size / 1024 / 1024) * 100 + '') / 100 + ' MB'
  } else if (size < (1024 * 1024 * 1024 * 1024)) {
    result = parseInt((size / 1024 / 1024 / 1024) * 100 + '') / 100 + ' GB'
  } else {
    result = parseInt((size / 1024 / 1024 / 1024 / 1024) * 100 + '') / 100 + ' TB'
  }
  if (flag) {
    result = result.split(' ').join('<span class="unit">') + '</span>'
  }
  return result
})
