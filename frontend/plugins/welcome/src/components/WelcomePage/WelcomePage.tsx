import React, { useEffect, FC, useState } from 'react';
import { makeStyles, createStyles } from '@material-ui/core/styles';
import { Theme } from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';
import {
  Container,
  FormControl,
  InputLabel,
  Select,
  Button,
  Grid,
  TextField,
} from '@material-ui/core';
import MenuItem from '@material-ui/core/MenuItem';
import { DefaultApi } from '../../api/apis';
import { EntUnit } from '../../api/models/EntUnit';
import { EntForm } from '../../api/models/EntForm';
import { EntVolume } from '../../api/models/EntVolume';
import { EntUser } from '../../api/models/EntUser';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import Swal from 'sweetalert2';
// import { EntDrug } from '../../api/models/EntDrug';

// header css
const HeaderCustom = {
  minHeight: '50px',
};

// css style
const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      flexGrow: 1,
    },

    paper: {
      marginTop: theme.spacing(2),
      marginBottom: theme.spacing(2),
    },
    formControl: {
      width: 300,
    },
    selectEmpty: {
      marginTop: theme.spacing(2),
    },
    container: {
      display: 'flex',
      flexWrap: 'wrap',
    },
    textField: {
      width: 300,
    },
  }),
);

interface drug {
  Unit: number;
  Form: number;
  Volume: number;
  User: number
  Strength: number;
  DrugType: string;
  Information: string;
}

const Drug: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();

  const [drug, setDrug] = React.useState<Partial<drug>>({});
  const [units, setUnits] = React.useState<EntUnit[]>([]);
  const [forms, setForms] = React.useState<EntForm[]>([]);
  const [users, setUsers] = React.useState<EntUser[]>([]);
  const [volumes, setVolumes] = React.useState<EntVolume[]>([]);

  const Toast = Swal.mixin({
    position: 'center',
    showConfirmButton: false,
    didOpen: toast => {
      toast.addEventListener('mouseenter', Swal.stopTimer);
      toast.addEventListener('mouseleave', Swal.resumeTimer);
    },
  });


  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>,) => {
    const name = event.target.name as keyof typeof Drug;
    const { value } = event.target;
    setDrug({ ...drug, [name]: value });
    console.log(drug);
  };

  const handleChangeNumber = (
    event: React.ChangeEvent<{ name?: string; value: number }>,) => {
    const name = event.target.name as keyof typeof Drug;
    const { value } = event.target;
    setDrug({ ...drug, [name]: +value });
    console.log(drug);
  };

  const getUsers = async () => {
    const res = await http.listUser({ limit: 10, offset: 0 });
    setUsers(res);
  };
  const getUnits = async () => {
    const res = await http.listUnit({ limit: 10, offset: 0 });
    setUnits(res);
  };
  const getForms = async () => {
    const res = await http.listForm({ limit: 10, offset: 0 });
    setForms(res);
  };
  const getVolumes = async () => {
    const res = await http.listVolume({ limit: 10, offset: 0 });
    setVolumes(res);
  };

  // Lifecycle Hooks
  useEffect(() => {
    getUsers();
    getUnits();
    getForms();
    getVolumes();
  }, []);

  function clear() {
    setDrug({});
  }

  function save() {
    const apiUrl = 'http://localhost:8080/api/v1/drugs';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(drug),

    };
    console.log(drug); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        if (data.status === true) {
          // clear();
          Toast.fire({
            icon: 'success',
            title: 'บันทึกข้อมูลสำเร็จ',
          });
        } else {
          Toast.fire({
            icon: 'error',
            title: 'บันทึกข้อมูลไม่สำเร็จ',
          });
        }
      });
  }

  return (
    <Page theme={pageTheme.home}>
      <Header style={HeaderCustom} title={`ระบบข้อมูลยา`}>
      </Header>
      <Content>
        <div style={{ textAlign: "center" }}>
          <Container maxWidth="sm">
            <Typography variant="h5" className={classes.formControl} style={{ marginLeft: 155 }}>
              ระบบข้อมูลยา
        </Typography>
            <Grid container spacing={6}>
              <Grid item xs={12}></Grid>
              <Grid item xs={6}>
                <FormControl
                  fullWidth
                  className={classes.formControl}
                  variant="outlined"
                >
                  <TextField
                    name="DrugType"
                    label="ยา"
                    variant="outlined"
                    type="string"
                    size="medium"
                    value={drug.DrugType || '' }
                    onChange={handleChange}
                  />
                </FormControl>
              </Grid>

              <Grid item xs={6}>
                <FormControl variant="outlined" className={classes.formControl} style={{ marginLeft: 20 }}>
                  <InputLabel>หน่วยนับ</InputLabel>
                  <Select
                    name="Unit"
                    value={drug.Unit || ''}
                    onChange={handleChange}
                    label="หน่วยนับ"

                  >
                    {units.map(item => {
                      return (
                        <MenuItem key={item.id} value={item.id}>{item.unitType}</MenuItem>
                      );
                    })}

                  </Select>
                </FormControl>
              </Grid>

              <Grid item xs={6}>
                <FormControl variant="outlined" className={classes.formControl} >
                  <InputLabel>รูปแบบยา</InputLabel>
                  <Select
                    name="Form"
                    value={drug.Form || ''}
                    onChange={handleChange}
                    label="รูปแบบยา"
                  >
                    {forms.map(item => {
                      return (
                        <MenuItem key={item.id} value={item.id}>{item.formType}</MenuItem>
                      );
                    })}
                  </Select>
                </FormControl>
              </Grid>

              <Grid item xs={3}>
                <FormControl
                  fullWidth
                  className={classes.formControl}
                  variant="outlined"
                  style={{ marginLeft: 20 }}
                >
                  <TextField
                    name="Strength"
                    label="ความแรง"
                    variant="outlined"
                    type="number"
                    size="medium"
                    value={drug.Strength || ''}
                    onChange={handleChangeNumber}
                  />
                </FormControl>
              </Grid>

              <Grid item xs={6}>
                <FormControl variant="outlined" className={classes.formControl}>
                  <InputLabel>หน่วย</InputLabel>
                  <Select
                    name="Volume"
                    value={drug.Volume || ''}
                    onChange={handleChange}
                    label="หน่วย"
                  >
                    {volumes.map(item => {
                      return (
                        <MenuItem key={item.id} value={item.id}>{item.volumeType}</MenuItem>
                      );
                    })}
                  </Select>
                </FormControl>

                <FormControl variant="outlined" className={classes.formControl} style={{ marginTop: 20 }}>
                  <InputLabel>ผู้บันทึก</InputLabel>
                  <Select
                    name="User"
                    value={drug.User || ''}
                    onChange={handleChange}
                    label="ผู้บันทึก"
                  >
                    {users.map(item => {
                      return (
                        <MenuItem key={item.id} value={item.id}>{item.username}</MenuItem>
                      );
                    })}
                  </Select>
                </FormControl>
              </Grid>

              <Grid item xs={6}>
                <FormControl variant="outlined" className={classes.formControl} style={{ marginLeft: 20 }}>
                  <TextField
                    name="Information"
                    value={drug.Information || ''}
                    label="ข้อมูลยา"
                    multiline
                    rows={5}
                    // defaultValue="ไม่มีข้อมูล"
                    variant="outlined"
                    onChange={handleChange}
                  />
                </FormControl>
              </Grid>

              <Grid item xs={2}></Grid>
              <Grid item xs={9}>
                <Button
                  variant="contained"
                  color="primary"
                  size="large"
                  startIcon={<SaveIcon />}
                  onClick = {save}
                  style={{ marginLeft: 25.4 }}
                >
                  บันทึก
              </Button>
                <Button
                  variant="contained"
                  color="secondary"
                  size="large"
                  startIcon={<SaveIcon />}
                  onClick={clear}
                  style={{ marginLeft: 30 }}
                >
                  ยกเลิก
              </Button>
              </Grid>
            </Grid>
          </Container>
        </div>
      </Content>
    </Page>
  );
};

export default Drug;

