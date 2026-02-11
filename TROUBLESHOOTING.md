# 🔧 Troubleshooting Guide

Common issues and solutions for Groupie Tracker v2 development.

---

## 🐳 Database Connection Issues

### **Problem: "Connection Refused" or "Failed to Ping Database"**

**Symptoms:**
```
❌ Database connection failed: dial tcp 127.0.0.1:5432: connect: connection refused
```

**Solution:**

```bash
# 1. Check if Docker is running
docker ps

# 2. If container is "Restarting" or not listed, check logs
docker compose logs postgres | tail -30

# 3. Fix based on error (see below)
```

---

### **Problem: Permission Denied Errors**

**Symptoms in logs:**
```
mkdir: can't create directory '/var/lib/postgresql/data/pgdata': Permission denied
chmod: /var/lib/postgresql/data/pgdata: Operation not permitted
```

**Solution:**

```bash
# Stop Docker
make docker-stop

# Fix permissions
sudo chown -R $(id -u):$(id -g) postgres-data/

# OR remove folder completely (fresh start)
sudo rm -rf postgres-data/

# Restart Docker
make docker-start

# Wait for database to be ready
sleep 10

# Verify it's healthy
docker ps
```

**Expected output:**
```
STATUS
Up X seconds (healthy)
```

---

### **Problem: Port 5432 Already in Use**

**Symptoms:**
```
Error: failed to bind host port 0.0.0.0:5432/tcp: address already in use
```

**Cause:** Local PostgreSQL service is running and using port 5432.

**Solution:**

```bash
# Check if PostgreSQL is running locally
sudo systemctl status postgresql

# If active, stop it
sudo systemctl stop postgresql

# Disable auto-start on boot
sudo systemctl disable postgresql

# Now start Docker
make docker-start
```

---

### **Problem: Database Tables Don't Exist**

**Symptoms:**
```
ERROR: relation "users" does not exist
```

**Cause:** Database is fresh/empty (migrations haven't run).

**Solution:**

```bash
# Stop Docker
make docker-stop

# Remove database
sudo rm -rf postgres-data/

# Start fresh (migrations run automatically)
make docker-start

# Wait for initialization
sleep 10

# Register a new user at http://localhost:8080/register
```

---

## 🔑 Session Secret Issues

### **Problem: Empty or Invalid SESSION_SECRET**

**Symptoms:**
- Login works but immediately redirects back to login
- "Invalid session token signature" in logs

**Solution:**

```bash
# Generate new SESSION_SECRET
make gen-secret

# Restart app
make dev
```

---

## 🛠️ Development Tools Issues

### **Problem: Air Not Found**

**Symptoms:**
```
❌ Air not found. Run 'make install-tools' first
```

**Solution:**

```bash
# Install Air
make install-tools

# If still not found, add to PATH
export PATH=$PATH:$(go env GOPATH)/bin

# Add to ~/.bashrc for persistence
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bashrc
source ~/.bashrc
```

---

### **Problem: .env File Missing**

**Symptoms:**
```
❌ .env file not found
```

**Solution:**

```bash
# Create .env from template
make env

# Generate SESSION_SECRET
make gen-secret
```

---

## 🔄 Cross-Machine Issues

### **Problem: Works on Laptop but Not on PC (or vice versa)**

**Common causes:**
1. Different UID/GID values
2. Different Docker setups
3. Missing .env file

**Solution:**

```bash
# On the new machine:

# 1. Pull latest code
git pull origin v2-development

# 2. Complete setup
make setup
make gen-secret

# 3. Start Docker
make docker-start

# 4. If permission errors, fix ownership
sudo chown -R $(id -u):$(id -g) postgres-data/
# OR
sudo rm -rf postgres-data/
make docker-start

# 5. Start dev server
make dev
```

---

## 🧹 Nuclear Option: Complete Reset

If nothing works, start completely fresh:

```bash
# 1. Stop everything
make docker-stop

# 2. Clean up
sudo rm -rf postgres-data/
rm -rf tmp/
rm .env

# 3. Rebuild from scratch
make setup
make gen-secret
make docker-start
sleep 10
make dev
```

---

## 📋 Quick Diagnostic Commands

**Check Docker:**
```bash
docker ps                          # Are containers running?
docker compose logs postgres       # What errors are happening?
docker inspect groupie-tracker-db  # Detailed container info
```

**Check Database Connection:**
```bash
# Connect to database directly
docker exec -it groupie-tracker-db psql -U groupie_user -d groupie_tracker

# Check tables exist
docker exec -it groupie-tracker-db psql -U groupie_user -d groupie_tracker -c "\dt"

# Check users
docker exec -it groupie-tracker-db psql -U groupie_user -d groupie_tracker -c "SELECT username, email FROM users;"
```

**Check .env:**
```bash
# Verify SESSION_SECRET is set
grep SESSION_SECRET .env

# Should NOT be the placeholder
# ❌ SESSION_SECRET=your_session_secret
# ✅ SESSION_SECRET=a1b2c3d4e5f6... (64 hex characters)
```

---

## 🆘 Still Having Issues?

1. Run `make verify` to check your setup
2. Review recent commits for breaking changes

---

## 💡 Prevention Tips

**Always:**
- Use `make docker-start` (NOT `docker compose up -d`)
- Run `make verify` after cloning on a new machine
- Keep `.env` out of Git (already in `.gitignore`)
- Don't manually edit `postgres-data/` folder

**Before pushing code:**
- Test on both machines if possible
- Commit `.env.example` changes
- Update this guide if you discover new issues!