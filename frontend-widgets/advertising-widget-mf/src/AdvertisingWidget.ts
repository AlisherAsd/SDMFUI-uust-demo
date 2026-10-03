import { defineComponent, h } from "vue";

export const htmlContent = `
    <img style="width: 100%; height: 200px" src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR-rSLeZKd7EPxtx6lbzxF2iej1uBCCTFRjVmWDxsI0jA&s=10" />
`;

export default defineComponent({
  name: "AdvertisingWidget",
  setup() {
    return () => h("div", { innerHTML: htmlContent });
  },
});
