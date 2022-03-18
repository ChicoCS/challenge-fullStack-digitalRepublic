import React from "react";
import { useNavigate } from "react-router-dom";
import { observer } from "mobx-react";
import PropTypes from "prop-types";
import { withStore } from "../../stores/context";

import Button from "@mui/material/Button";
import Card from "@mui/material/Card";
import CardContent from "@mui/material/CardContent";
import CardHeader from "@mui/material/CardHeader";
import Grid from "@mui/material/Grid";
import Typography from "@mui/material/Typography";

MainPageResult.propTypes = {
  mainPageStore: PropTypes.object.isRequired,
};

function MainPageResult(props) {
  const { mainPageStore } = props;

  let navigate = useNavigate();

  const backToMainPage = () => {
    mainPageStore.reset();
    navigate("/");
  };

  return (
    <div style={{ marginRight: "10%", marginLeft: "10%" }}>
      <Card style={{ backgroundColor: "#98cdff" }}>
        <CardHeader
          style={{ color: "#3a84ff" }}
          title="Result of the calculation"
        />
        <Card>
          <CardContent>
            <Grid container spacing={2}>
              {mainPageStore.calculationResult.qty_can18l > 0 && (
                <Grid item xs={12}>
                  <Typography variant="h5">{`18L can = ${mainPageStore.calculationResult.qty_can18l}`}</Typography>
                </Grid>
              )}
              {mainPageStore.calculationResult.qty_can3_6l > 0 && (
                <Grid item xs={12}>
                  <Typography variant="h5">{`3,6L can = ${mainPageStore.calculationResult.qty_can3_6l}`}</Typography>
                </Grid>
              )}
              {mainPageStore.calculationResult.qty_can2_5l > 0 && (
                <Grid item xs={12}>
                  <Typography variant="h5">{`2,5L can = ${mainPageStore.calculationResult.qty_can2_5l}`}</Typography>
                </Grid>
              )}
              {mainPageStore.calculationResult.qty_can0_5l > 0 && (
                <Grid item xs={12}>
                  <Typography variant="h5">{`0,5L can = ${mainPageStore.calculationResult.qty_can0_5l}`}</Typography>
                </Grid>
              )}
              {mainPageStore.calculationResult.liters > 0 && (
                <Grid item xs={12}>
                  <Typography variant="h5">{`Total liters of paint = ${mainPageStore.calculationResult.total}L`}</Typography>
                </Grid>
              )}
            </Grid>
          </CardContent>
        </Card>
        <div style={divButton}>
          <Button
            style={{ marginRight: 10 }}
            color="primary"
            variant="contained"
            onClick={backToMainPage}
          >
            Back
          </Button>
        </div>
      </Card>
    </div>
  );
}

const divButton = {
  height: 50,
  display: "flex",
  justifyContent: "flex-end",
  alignItems: "center",
};

export default withStore(observer(MainPageResult));
