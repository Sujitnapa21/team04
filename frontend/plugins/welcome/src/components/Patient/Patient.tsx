import React, { FC, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import MenuIcon from '@material-ui/icons/Menu';
import AccountCircle from '@material-ui/icons/AccountCircle';
import Link from '@material-ui/core/Link';
import Swal from 'sweetalert2'; // alert
import {
  Menu,
  MenuProps,
  AppBar,
  Toolbar,
  Typography,
  IconButton,
  ListItemIcon,
  ListItemText,
  Grid,
  TextField,
  Container,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
  Button
} from '@material-ui/core';
import LocalHotelRoundedIcon from '@material-ui/icons/LocalHotelRounded';
import { withStyles } from '@material-ui/core/styles';
import SearchIcon from '@material-ui/icons/Search';

// css style 
const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1,
    marginTop: theme.spacing(10),
  },
  textField: {
    marginLeft: theme.spacing(1),
    marginRight: theme.spacing(1),
    width: '25ch',
    color: "blue"
  },
  menuButton: {
    marginRight: theme.spacing(2),
  },
  title: {
    flexGrow: 1,
  },
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  paper: {
    marginTop: theme.spacing(2),
    marginBottom: theme.spacing(2),
  },
  formControl: {
    width: 455,
  },
  selectEmpty: {
    marginTop: theme.spacing(2),
  },

  header: {
    textAlign: 'center'
  },

  buttonSty: {
    margin: 'auto',
    display: 'block',
    maxWidth: '100%',
    maxHeight: '100%',
    marginBottom: 50
  },
  logoutButton: {
    marginLeft: 10,
    marginRight: 10,
    color: 'white'
  },

}));
const StyledMenu = withStyles({
  paper: {
    border: '1px solid #d3d4d5',
  },
})((props: MenuProps) => (
  <Menu
    elevation={0}
    getContentAnchorEl={null}
    anchorOrigin={{
      vertical: 'bottom',
      horizontal: 'center',
    }}
    transformOrigin={{
      vertical: 'top',
      horizontal: 'center',
    }}
    {...props}
  />
));

const StyledMenuItem = withStyles((theme) => ({
  root: {
    '&:focus': {
      backgroundColor: theme.palette.primary.main,
      '& .MuiListItemIcon-root, & .MuiListItemText-primary': {
        color: theme.palette.common.white,
      },
    },
  },
}))(MenuItem);




import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command
import { EntEmployee } from '../../api/models/EntEmployee'; //import interface Employee 
import { EntCategory } from '../../api/models/EntCategory'; //import interface Category
import { EntNametitle } from '../../api/models/EntNametitle'; //import interface Nametitle
import { EntBloodtype } from '../../api/models/EntBloodtype'; //import interface Bloodtype
import { EntGender } from '../../api/models/EntGender'; //import interface Gender  

interface Patient {
  Idcard: string;
  Category: number;
  Nametitle: number;
  PatientName: string;
  Bloodtype: number;
  Gender: number;
  Address: string;
  Congenital: string;
  Allergic: string;
  Employee: number;
}


const Patient: FC<{}> = () => {
  const classes = useStyles();
  const api = new DefaultApi();

  const [showInputError, setShowInputError] = React.useState(false); // for error input show
  const [patient, setPatient] = React.useState<Partial<Patient>>({});
  const [employees, setEmployees] = React.useState<EntEmployee[]>([]);
  const [categorys, setCategorys] = React.useState<EntCategory[]>([]);
  const [nametitles, setNametitles] = React.useState<EntNametitle[]>([]);
  const [bloodtypes, setBloodtypes] = React.useState<EntBloodtype[]>([]);
  const [genders, setGenders] = React.useState<EntGender[]>([]);
  const [idcardError, setidcardError] = React.useState('');
  const [congenitalError, setcongenitalError] = React.useState('');
  const [allergicError, setallergicError] = React.useState('');

  const [anchorEl, setAnchorEl] = React.useState<null | HTMLElement>(null);

  const handleClick = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorEl(event.currentTarget);
  };

  const handleClose = () => {
    setAnchorEl(null);
  };

  function redirectToPatient() {
    window.location.href = "http://localhost:3000/patient"
}

function redirectToSearchPatient() {
  window.location.href = "http://localhost:3000/searchpatient"
}

  // alert setting
  const Toast = Swal.mixin({
    toast: true,
    position: undefined,
    showConfirmButton: false,
    timer: 3000,
    timerProgressBar: true,
    didOpen: toast => {
      toast.addEventListener('mouseenter', Swal.stopTimer);
      toast.addEventListener('mouseleave', Swal.resumeTimer);
    },
  });

  const getEmployee = async () => {
    const res = await api.listEmployee({ limit: undefined, offset: 0 });
    setEmployees(res);
  };

  const getCategory = async () => {
    const res = await api.listCategory({ limit: 4, offset: 0 });
    setCategorys(res);
  };

  const getNametitle = async () => {
    const res = await api.listNametitle({ limit: 10, offset: 0 });
    setNametitles(res);
  };

  const getBloodtype = async () => {
    const res = await api.listBloodtype({ limit: 4, offset: 0 });
    setBloodtypes(res);
  };
  const getGender = async () => {
    const res = await api.listGender({ limit: 2, offset: 0 });
    setGenders(res);
  };

  // Lifecycle Hooks
  useEffect(() => {
    getEmployee();
    getCategory();
    getNametitle();
    getBloodtype();
    getGender();

  }, []);

  // set data to object patient
  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: any }>,
  ) => {
    const name = event.target.name as keyof typeof patient;
    const { value } = event.target;
    const validateValue = value.toString()
    setPatient({ ...patient, [name]: value });
    console.log(patient);
    checkPattern(name, validateValue);

  };
  // ฟังก์ชั่นสำหรับ validate เลขประจำตัวประชาชน
   const Idcard = (val: string) => {
    return val.length == 13 ? true : false;
  }

  // ฟังก์ชั่นสำหรับ validate โรคประจำตัว
  const Congenital = (val: string) => {
    return val.match("[โรค, ไม่มี]");
  }

  // ฟังก์ชั่นสำหรับ validate ประวัติการแพ้ยา
  const Allergic = (val: string) => {
    return val.match("[ตัวยาชื่อ, ไม่มี, ยา]");
  }



  // สำหรับตรวจสอบรูปแบบข้อมูลที่กรอก ว่าเป็นไปตามที่กำหนดหรือไม่
  const checkPattern  = (id: string, value: string) => {
    switch(id) {
      case 'Idcard':
        Idcard(value) ? setidcardError('') : setidcardError('เลขประจำตัวประชาชน 13 หลัก');
        return;
        case 'Congenital':
          Congenital(value) ? setcongenitalError('') : setcongenitalError('ขึ้นต้นด้วย โรค, ไม่มี,-');
          return;
          case 'Allergic':
            Allergic(value) ? setallergicError('') : setallergicError('ขึ้นต้นด้วย ตัวยาชื่อ, ไม่มี, -, ยา');
            return;
      default:
        return;
    }
  }
  const alertMessage = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
  }

  const checkCaseSaveError = (field: string) => {
    switch(field) {
      case 'Idcard':
        alertMessage("error","เลขประจำตัวประชาชน 13 หลัก");
        return;
      case 'Congenital':
        alertMessage("error","โรคประจำตัวต้องขึ้นต้นด้วย โรค, ไม่มี,-");
        return;
      case 'Allergic':
        alertMessage("error","ประวัติแพ้ยาต้องขึ้นต้นด้วย ตัวยาชื่อ, ไม่มี, -, ยา");
        return;
      default:
        alertMessage("error","บันทึกข้อมูลไม่สำเร็จ");
        return;
    }
  }

   // clear input form
  function clear() {
   setPatient({});
   setShowInputError(false);
  }

  // function save data
  function save() {
    // setShowInputError(true)
    // let { Idcard, PatientName, Address, Congenital, Allergic } = Patient;
    // if (!Idcard || !PatientName || !Address || !Congenital || !Allergic) {
    //   Toast.fire({
    //     icon: 'error',
    //     title: 'กรุณากรอกข้อมูลให้ครบถ้วน',
    //   });
    //   return;
    // }

       
//เช็คแล้วเก็บค่าไว้ใน employee
patient.Employee = employees.filter(emp => emp.userId === window.localStorage.getItem("username"))[0].id;


    const apiUrl = 'http://localhost:8080/api/v1/patients';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(patient),
    };

    console.log(patient); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {console.log(data.save)
        console.log(requestOptions)
        if (data.status == true) {
          clear();
          Toast.fire({
            icon: 'success',
            title: 'บันทึกข้อมูลสำเร็จ',
          });
        } else {
          checkCaseSaveError(data.error.Name)
        }
      });
};
  //Java 
  function redirecLogOut() {
    //redirec Page ... http://localhost:3000/
    window.location.href = "http://localhost:3000/";
  }

  return (
    <div className={classes.root}>
       <AppBar position="fixed" >
        <Toolbar>
          <IconButton 
        onClick={handleClick}>
              <MenuIcon />
          </IconButton>
          <StyledMenu
        id="customized-menu"
        anchorEl={anchorEl}
        keepMounted
        open={Boolean(anchorEl)}
        onClose={handleClose}
      >

        <StyledMenuItem button onClick={redirectToPatient}>
          <ListItemIcon>
            <LocalHotelRoundedIcon fontSize="default" />
          </ListItemIcon>
          <ListItemText primary="Add Patient" />
        </StyledMenuItem>

        <StyledMenuItem button onClick={redirectToSearchPatient}>
          <ListItemIcon>
            <SearchIcon fontSize="default" />
          </ListItemIcon>
          <ListItemText primary="Search Patient" />
        </StyledMenuItem>
      </StyledMenu>

          <Typography variant="h4" className={classes.title}>
            ระบบจัดการโรคติดต่อ
          </Typography>
            <IconButton
              aria-label="account of current user"
              aria-controls="menu-appbar"
              aria-haspopup="true"
              color="inherit"
            >
              <AccountCircle />
              <Typography>
                <Link variant="h6" onClick={redirecLogOut} className={classes.logoutButton}>
                  LOGOUT
                </Link>
              </Typography>
            </IconButton>
        </Toolbar>
      </AppBar>
      <Container maxWidth="sm">

        <Grid container spacing={3}>

          <Grid item xs={10}>
            <h2 style={{ textAlign: 'center' }} >ข้อมูลประจำตัวผู้ป่วย </h2>
          </Grid>

          <Grid item xs={10}>
          <TextField
                error = {idcardError ? true : false}
                className={classes.formControl}
                name="Idcard"
                label="เลขประจำตัวประชาชน"
                variant="outlined"
                type="string"
                inputProps={{ maxLength: 13 }}
                helperText= {idcardError}
                value={patient.Idcard || ''}
                onChange={handleChange}
              />
            </Grid>

          <Grid item xs={12}>
            <FormControl variant="outlined" className={classes.formControl}>
              <InputLabel >ประเภทผู้ป่วย</InputLabel>
              <Select
                name="Category"
                error={!patient.Category && showInputError}
                value={patient.Category || ''}
                onChange={handleChange}
                label="ประเภทผู้ป่วย"
              >
                {categorys.map(item => {
                  return (
                    <MenuItem key={item.id} value={item.id}>
                      {item.categoryName}
                    </MenuItem>
                  );
                })}
              </Select>
            </FormControl>
          </Grid>

          <Grid item xs={12}>
            <FormControl variant="outlined" className={classes.formControl}>
              <InputLabel >คำนำหน้าชื่อ</InputLabel>
              <Select
                name="Nametitle"
                error={!patient.Nametitle && showInputError}
                value={patient.Nametitle || ''}
                onChange={handleChange}
                label="คำนำหน้าชื่อ"
                fullWidth
              >
                {nametitles.map(item => {
                  return (
                    <MenuItem key={item.id} value={item.id}>
                      {item.title}
                    </MenuItem>
                  );
                })}
              </Select>
            </FormControl>
          </Grid>

          <Grid item xs={10}>
            <TextField
              required={true}
              error={!patient.PatientName && showInputError}
              name="PatientName"
              label="ชื่อ-นามสกุล"
              variant="outlined"
              fullWidth
              multiline
              value={patient.PatientName || ""}
              onChange={handleChange}
            />
          </Grid>

          <Grid item xs={12}>
            <FormControl variant="outlined" className={classes.formControl}>
              <InputLabel>เพศ</InputLabel>
              <Select
                name="Gender"
                error={!patient.Gender && showInputError}
                value={patient.Gender || ''}
                onChange={handleChange}
                label="เพศ"
              >
                {genders.map(item => {
                  return (
                    <MenuItem key={item.id} value={item.id}>
                      {item.genderName}
                    </MenuItem>
                  );
                })}
              </Select>
            </FormControl>
          </Grid>

          <Grid item xs={12}>
            <FormControl variant="outlined" className={classes.formControl}>
              <InputLabel >กรุ๊ปเลือด</InputLabel>
              <Select
                name="Bloodtype"
                error={!patient.Bloodtype && showInputError}
                value={patient.Bloodtype || ''}
                onChange={handleChange}
                label="กรุ๊ปเลือด"
                fullWidth
              >
                {bloodtypes.map(item => {
                  return (
                    <MenuItem key={item.id} value={item.id}>
                      {item.bloodtypeName}
                    </MenuItem>
                  );
                })}
              </Select>
            </FormControl>
          </Grid>

          <Grid item xs={10}>
            <TextField
              name="Address"
              error={!patient.Address && showInputError}
              label="ที่อยู่"
              variant="outlined"
              fullWidth
              multiline
              value={patient.Address || ''}
              onChange={handleChange}
            />
          </Grid>

          <Grid item xs={10}>
            <h2 style={{ textAlign: 'center' }}> ข้อมูลทางการแพทย์ </h2>
          </Grid>

          <Grid item xs={10}>
          <TextField
                error = {congenitalError ? true : false}
                className={classes.formControl}
                name="Congenital"
                label="โรคประจำตัว"
                helperText= {congenitalError}
                value={patient.Congenital || ''}
                variant="outlined"
                onChange={handleChange}
              />
            </Grid>

          <Grid item xs={10}>
          <TextField
                error = {allergicError ? true : false}
                className={classes.formControl}
                name="Allergic"
                label="ประวัติแพ้ยา"
                helperText= {allergicError}
                value={patient.Allergic || ''}
                variant="outlined"
                onChange={handleChange}
              />
            </Grid>


          <Grid item xs={10}>
              <TextField
                required={true}
                disabled // ห้ามแก้ไข
                // id="name"
                name="name"
                type="string"
                label="รหัส"
                variant="outlined"
                fullWidth
                multiline
                value={ window.localStorage.getItem("username") || ""}
                onChange={handleChange}
              />
            </Grid>
{/* 

            <Grid item xs={12}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel >รหัสพนักงาน</InputLabel>
                <Select
                  name="employee"
                  value={patient.employee  || ''}
                  onChange={handleChange}
                  label="รหัสพนักงาน"
                >
                  {employees.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.userId}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
          </Grid> */}

          <Grid item xs={10}>
            <Button
              name="saveData"
              size="large"
              variant="contained"
              color="primary"
              disableElevation
              className={classes.buttonSty}
              onClick={save}
            >
              บันทึกข้อมูล
              </Button>
          </Grid>
        </Grid>
      </Container>
    </div>
  );
};

export default Patient;