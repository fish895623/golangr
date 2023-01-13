window.onload = () => {
  get();
};

var myEditor;
const editor = document.querySelector("#classic");

BalloonEditor.create(editor)
  .then((editor) => {
    myEditor = editor;
  })
  .catch((error) => {
    console.error(error);
  });
function send() {
  axios({
    url: "/api",
    method: "POST",
    data: {
      load: false,
      data: myEditor.getData(),
    },
  });
}
function get() {
  axios({
    url: "/api",
    method: "POST",
    data: {
      load: true,
    },
  }).then((req) => {
    myEditor.setData(req.data.data);
  });
}
