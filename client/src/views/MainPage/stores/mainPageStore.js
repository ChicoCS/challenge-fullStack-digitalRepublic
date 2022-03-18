import { action, makeObservable, observable, computed } from "mobx";
import Services from "../../../services/Services";

class MainPageStore {
  walls = [
    {
      name: "A",
      height: "0",
      width: "0",
      n_doors: "0",
      n_windows: "0",
    },
    {
      name: "B",
      height: "0",
      width: "0",
      n_doors: "0",
      n_windows: "0",
    },
    {
      name: "C",
      height: "0",
      width: "0",
      n_doors: "0",
      n_windows: "0",
    },
    {
      name: "D",
      height: "0",
      width: "0",
      n_doors: "0",
      n_windows: "0",
    },
  ];

  calculationResult = {};

  fetching = false;

  reset() {
    this.walls = [
      {
        name: "A",
        height: "0",
        width: "0",
        n_doors: "0",
        n_windows: "0",
      },
      {
        name: "B",
        height: "0",
        width: "0",
        n_doors: "0",
        n_windows: "0",
      },
      {
        name: "C",
        height: "0",
        width: "0",
        n_doors: "0",
        n_windows: "0",
      },
      {
        name: "D",
        height: "0",
        width: "0",
        n_doors: "0",
        n_windows: "0",
      },
    ];
    this.calculationResult = {};
  }

  handleChangeWall(target, index) {
    const { name, value } = target;
    this.walls[index][name] = value;
  }

  validateFields(data) {
    let validation = 0;
    data.forEach((wall) => {
      if (
        wall.height > parseFloat(0) &&
        wall.width > parseFloat(0) &&
        wall.qty_doors >= parseInt(0) &&
        wall.qty_windows >= parseInt(0)
      ) {
        validation += 1;
      }
    });

    return validation === 4;
  }

  async calculate(navigate) {
    try {
      this.fetching = true;

      const data = [];

      this.walls.map((wall) =>
        data.push({
          height: parseFloat(wall.height),
          width: parseFloat(wall.width),
          qty_doors: wall.n_doors ? parseInt(wall.n_doors) : parseInt(0),
          qty_windows: wall.n_windows ? parseInt(wall.n_windows) : parseInt(0),
        })
      );

      const validateData = this.validateFields(data);

      if (validateData) {
        const response = await Services.Calculate(data);

        if (response.error) {
          alert(`${response.error}`);
        }

        if (response.data) {
          this.calculationResult = response.data

          navigate("/paint/result");
        }
      }

      if (!validateData) {
        alert("Invalid fields.");
      }
    } finally {
      this.fetching = false;
    }
  }

  get enableButtonCalculate() {
    let validation = 0;
    this.walls.forEach((wall) => {
      if (wall.height > 0 && wall.width > 0) {
        validation++;
      }
    });

    return validation === 4;
  }

  constructor() {
    makeObservable(this, {
      walls: observable,
      fetching: observable,
      calculationResult: observable,
      handleChangeWall: action.bound,
      calculate: action.bound,
      validateFields: action.bound,
      reset: action.bound,
      enableButtonCalculate: computed,
    });
  }
}

export default MainPageStore;
