import { action, makeObservable, observable, computed } from "mobx";
import Services from "../../../services/Services";

class MainPageStore {
  walls = [
    {
      name: "A",
      height: null,
      width: null,
      n_doors: null,
      n_windows: null,
    },
    {
      name: "B",
      height: null,
      width: null,
      n_doors: null,
      n_windows: null,
    },
    {
      name: "C",
      height: null,
      width: null,
      n_doors: null,
      n_windows: null,
    },
    {
      name: "D",
      height: null,
      width: null,
      n_doors: null,
      n_windows: null,
    },
  ];

  fetching = false;

  handleChangeWall(target, index) {
    const { name, value } = target;
    this.walls[index][name] = value;
  }

  validateFields(data) {
    let validation = false;
    data.forEach((wall) => {
      if (
        wall.height > parseFloat(0) &&
        wall.width > parseFloat(0) &&
        wall.n_doors >= parseInt(0) &&
        wall.n_windows >= parseInt(0)
      ) {
        validation = true;
      } else {
        validation = false;
      }
    });
    return validation;
  }

  async calculate() {
    try {
      this.fetching = true;

      const data = [];

      this.walls.map((wall) =>
        data.push({
          height: parseFloat(wall.height),
          width: parseFloat(wall.width),
          n_doors: wall.n_doors ? parseInt(wall.n_doors) : parseInt(0),
          n_windows: wall.n_windows ? parseInt(wall.n_windows) : parseInt(0),
        })
      );

      const negativeNumberExists = this.validateFields(data);

      if (negativeNumberExists) {
        const response = await Services.Calculate(data);
        if (response.error) {
          alert(`${response.error}`);
        }
      } else {
        alert("Invalid fields.");
      }
    } finally {
      this.fetching = false;
    }
  }

  get enableButtonCalculate() {
    let validation = false;
    this.walls.forEach((wall) => {
      if (wall.height == null || wall.width == null) {
        validation = true;
      } else {
        validation = false;
      }
    });
    return validation;
  }

  constructor() {
    makeObservable(this, {
      walls: observable,
      fetching: observable,
      handleChangeWall: action.bound,
      calculate: action.bound,
      validateFields: action.bound,
      enableButtonCalculate: computed,
    });
  }
}

export default MainPageStore;
