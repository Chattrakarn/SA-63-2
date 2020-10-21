import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import { EntDrug } from '../../api/models/EntDrug';

const useStyles = makeStyles({
 table: {
   minWidth: 650,
 },
});
 
export default function ComponentsTable() {
 const classes = useStyles();
 const api = new DefaultApi();
 const [drugs, setDrugs] = useState<EntDrug[]>([]);
 const [loading, setLoading] = useState(true);
 
 useEffect(() => {
   const getDrugs = async () => {
     const res = await api.listDrug({ limit: 10, offset: 0 });
     setLoading(false);
     setDrugs(res);
   };
   getDrugs();
 }, [loading]);
 
 const deleteDrugs = async (id: number) => {
   const res = await api.deleteDrug({ id: id });
   setLoading(true);
 };
 
 return (
   <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
           <TableCell align="center">No.</TableCell>
           <TableCell align="center">DrugType</TableCell>
           <TableCell align="center">UnitType</TableCell>
           <TableCell align="center">Form</TableCell>
           <TableCell align="center">Strength</TableCell>
           <TableCell align="center">Volume</TableCell>
           <TableCell align="center">Information</TableCell>
           <TableCell align="center">Username</TableCell>
           <TableCell align="center">Manage</TableCell>
         </TableRow>
       </TableHead>
       <TableBody>
         {drugs.map(item => (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{item.drugType}</TableCell>
             <TableCell align="center">{item.edges?.unit?.unitType}</TableCell>
             <TableCell align="center">{item.edges?.form?.formType}</TableCell>
             <TableCell align="center">{item.strength}</TableCell>
             <TableCell align="center">{item.edges?.volume?.volumeType}</TableCell>
             <TableCell align="center">{item.information}</TableCell>
             <TableCell align="center">{item.edges?.user?.username}</TableCell>
             <TableCell align="center">
               <Button
                 onClick={() => {
                  deleteDrugs(item.id);
                 }}
                 style={{ marginLeft: 10 }}
                 variant="contained"
                 color="secondary"
               >
                 Delete
               </Button>
             </TableCell>
           </TableRow>
         ))}
       </TableBody>
     </Table>
   </TableContainer>
 );
}