// 转换时间戳
export function formatDate(timestamp) {
  console.log(timestamp);
  const date = new Date(parseInt(timestamp)*1000);
  const year = date.getFullYear();
  const month = addLeadingZero(date.getMonth() + 1);
  const day = addLeadingZero(date.getDate());
  const hour = addLeadingZero(date.getHours());
  const minute = addLeadingZero(date.getMinutes());
  return `${year}-${month}-${day} ${hour}:${minute}`;
}

export function addLeadingZero(number) {
  return number < 10 ? `0${number}` : number;
}

