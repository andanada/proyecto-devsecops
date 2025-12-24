# Plan de Recuperación ante Desastres (DRP)
## Disaster Recovery Plan - Proyecto DevSecOps

---

## 1. Objetivos del DRP

### 1.1 Recovery Time Objective (RTO)
- **Aplicación Web:** 30 minutos
- **Pipeline CI/CD:** 1 hora
- **Repositorio GitHub:** 15 minutos (mediante mirrors)

### 1.2 Recovery Point Objective (RPO)
- **Código fuente:** 0 minutos (sincronización continua con GitHub)
- **Artifacts de build:** 24 horas
- **Logs de aplicación:** 1 hora
- **Base de datos (futuro):** 15 minutos

---

## 2. Estrategias de Backup

### 2.1 Repositorio de Código
**Ubicación:** GitHub (andanada/proyecto-devsecops)

**Backup primario:**
- GitHub mantiene múltiples réplicas geográficas
- Commits permanentes con historial completo

**Backup secundario:**
- Mirror en servidor local: `/home/grup1/proyecto-devsecops`
- Actualizaciones automáticas mediante git pull

**Frecuencia:** Continua (cada push)

**Comando de restauración:**
```bash
git clone https://github.com/andanada/proyecto-devsecops.git
