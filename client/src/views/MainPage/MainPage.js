import React, { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { observer } from "mobx-react";
import PropTypes from "prop-types";
import { withStore } from "../../stores/context";

import Button from "@mui/material/Button";
import Card from "@mui/material/Card";
import CardContent from "@mui/material/CardContent";
import CardHeader from "@mui/material/CardHeader";
import Table from "@mui/material/Table";
import TableBody from "@mui/material/TableBody";
import TableCell from "@mui/material/TableCell";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";
import TextField from "@mui/material/TextField";

MainPage.propTypes = {
  mainPageStore: PropTypes.object.isRequired,
};

function MainPage(props) {
  const { mainPageStore } = props;

  let navigate = useNavigate();

  const calculate = async () => {
    await mainPageStore.calculate(navigate);
  };

  useEffect(() => {
    mainPageStore.reset();
  }, [mainPageStore]);

  return (
    <div style={{ marginRight: "10%", marginLeft: "10%" }}>
      <Card style={{ backgroundColor: "#98cdff" }}>
        <CardHeader
          style={{ color: "#3a84ff" }}
          title="Calculate amount of paint required for room"
        />

        <Card>
          <CardContent>
            <Table size="small">
              <TableHead>
                <TableRow>
                  <TableCell style={headerTableStyle} align="left">
                    Wall
                  </TableCell>
                  <TableCell style={headerTableStyle} align="left">
                    Wall Height
                  </TableCell>
                  <TableCell style={headerTableStyle} align="left">
                    Wall Width
                  </TableCell>
                  <TableCell style={headerTableStyle} align="left">
                    Number of Doors
                  </TableCell>
                  <TableCell style={headerTableStyle} align="left">
                    Number of Windows
                  </TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {mainPageStore.walls.map((row, index) => (
                  <TableRow key={index}>
                    <TableCell style={headerTableStyle} align="center">
                      {row.name}
                    </TableCell>
                    <TableCell>
                      <TextField
                        required
                        name="height"
                        size="small"
                        variant="filled"
                        label="Height"
                        value={row.height}
                        type="number"
                        onChange={({ target }) =>
                          mainPageStore.handleChangeWall(target, index)
                        }
                        inputProps={inputPropsHeightWidth}
                      />
                    </TableCell>
                    <TableCell>
                      <TextField
                        required
                        name="width"
                        size="small"
                        variant="filled"
                        label="Width"
                        value={row.width}
                        type="number"
                        onChange={({ target }) =>
                          mainPageStore.handleChangeWall(target, index)
                        }
                        inputProps={inputPropsHeightWidth}
                      />
                    </TableCell>
                    <TableCell>
                      <TextField
                        name="n_doors"
                        size="small"
                        variant="filled"
                        label="Doors"
                        value={row.n_doors}
                        type="number"
                        onChange={({ target }) =>
                          mainPageStore.handleChangeWall(target, index)
                        }
                        inputProps={{ min: 0 }}
                      />
                    </TableCell>
                    <TableCell>
                      <TextField
                        name="n_windows"
                        size="small"
                        variant="filled"
                        label="Windows"
                        value={row.n_windows}
                        type="number"
                        onChange={({ target }) =>
                          mainPageStore.handleChangeWall(target, index)
                        }
                        inputProps={{ min: 0 }}
                      />
                    </TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </CardContent>
        </Card>
        <div style={divButton}>
          <Button
            disabled={!mainPageStore.enableButtonCalculate}
            style={{ marginRight: 10 }}
            color="primary"
            variant="contained"
            onClick={calculate}
          >
            Calculate
          </Button>
        </div>
      </Card>
    </div>
  );
}

const headerTableStyle = {
  color: "#3a84ff",
  fontWeight: "bold",
  fontSize: 20,
};

const inputPropsHeightWidth = {
  placeholder: "ex: 1,65",
  step: "0.05",
  min: 0,
};

const divButton = {
  height: 50,
  display: "flex",
  justifyContent: "flex-end",
  alignItems: "center",
};

export default withStore(observer(MainPage));
