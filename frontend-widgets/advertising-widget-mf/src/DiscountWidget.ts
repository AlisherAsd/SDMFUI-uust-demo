import { defineComponent, h } from "vue";

export const htmlContent = `
    <div style="width: 100%; height: 50px; background-color: blue; color: white" src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR-rSLeZKd7EPxtx6lbzxF2iej1uBCCTFRjVmWDxsI0jA&s=10">
        <b>ШОК!!!! Скидки до 0% !!!!! Никогда не было и вот опять!!!</b>
    </div>
`;

export default defineComponent({
  name: "DiscountWidget",
  setup() {
    return () => h("div", { innerHTML: htmlContent });
  },
});
